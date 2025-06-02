import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
  Container,
  Grid,
  Paper,
  NumberInput,
  Button,
  Group,
  Modal,
  Text,
  Stack,
  ActionIcon,
  Tooltip,
  AppShell,
  Title,
} from '@mantine/core';
import { DateTimePicker } from '@mantine/dates';
import { notifications } from '@mantine/notifications';
import { IconCalendar, IconMapPin, IconMail } from '@tabler/icons-react';
import { tablesApi, bookingApi } from '../services/api';
import type { Table } from '../types/api';
import dayjs from 'dayjs';
import { Footer } from '../components/Footer';
import { cafeConfig } from '../config/cafe';

interface TableLayoutProps {
  tables: Table[];
  onTableClick: (table: Table) => void;
  guestsCount: number | '';
}

function TableLayout({ tables, onTableClick, guestsCount }: TableLayoutProps) {
  const renderChairs = (
    tableWidth: number, 
    tableHeight: number, 
    seatsNumber: number, 
    angle: number, 
    isUnavailable: boolean
  ): React.ReactNode[] => {
    const chairs: React.ReactNode[] = [];
    const chairSize = Math.min(tableWidth, tableHeight) * 0.35;
    
    // Простое расположение стульев в зависимости от количества мест
    const positions = [];
    
    if (seatsNumber === 2) {
      // 2 стула напротив друг друга по длинной стороне
      positions.push(
        { x: 0, y: -tableHeight * 0.8, angle: 0 },  // сверху
        { x: 0, y: tableHeight * 0.8, angle: 180 }, // снизу
      );
    } else if (seatsNumber === 4) {
      // 4 стула по длинным сторонам
      positions.push(
        { x: -tableWidth * 0.2, y: -tableHeight * 0.8, angle: 0 },    // сверху слева
        { x: -tableWidth * 0.2, y: tableHeight * 0.8, angle: 180 },   // снизу слева
        { x: tableWidth * 0.2, y: -tableHeight * 0.8, angle: 0 },     // сверху справа
        { x: tableWidth * 0.2, y: tableHeight * 0.8, angle: 180 },    // снизу справа
      );
    } else if (seatsNumber === 6) {
      // 6 стульев: по 3 с каждой длинной стороны
      const xOffset = tableWidth * 0.25;
      positions.push(
        // Верхняя сторона
        { x: -xOffset, y: -tableHeight * 0.8, angle: 0 },
        { x: 0, y: -tableHeight * 0.8, angle: 0 },
        { x: xOffset, y: -tableHeight * 0.8, angle: 0 },
        // Нижняя сторона
        { x: -xOffset, y: tableHeight * 0.8, angle: 180 },
        { x: 0, y: tableHeight * 0.8, angle: 180 },
        { x: xOffset, y: tableHeight * 0.8, angle: 180 },
      );
    } else if (seatsNumber === 8) {
      // 8 стульев: по 4 с каждой длинной стороны
      const xOffset = tableWidth * 0.2;
      positions.push(
        // Верхняя сторона
        { x: -xOffset * 1.5, y: -tableHeight * 0.8, angle: 0 },
        { x: -xOffset * 0.5, y: -tableHeight * 0.8, angle: 0 },
        { x: xOffset * 0.5, y: -tableHeight * 0.8, angle: 0 },
        { x: xOffset * 1.5, y: -tableHeight * 0.8, angle: 0 },
        // Нижняя сторона
        { x: -xOffset * 1.5, y: tableHeight * 0.8, angle: 180 },
        { x: -xOffset * 0.5, y: tableHeight * 0.8, angle: 180 },
        { x: xOffset * 0.5, y: tableHeight * 0.8, angle: 180 },
        { x: xOffset * 1.5, y: tableHeight * 0.8, angle: 180 },
      );
    }

    // Создаем стулья на основе позиций с учетом поворота стола
    positions.forEach(({ x, y, angle: chairAngle }, i) => {
      // Поворачиваем координаты стула в соответствии с углом поворота стола
      const radians = (angle * Math.PI) / 180;
      const rotatedX = x * Math.cos(radians) - y * Math.sin(radians);
      const rotatedY = x * Math.sin(radians) + y * Math.cos(radians);

      chairs.push(
        <div
          key={i}
          style={{
            position: 'absolute',
            width: `${chairSize}px`,
            height: `${chairSize}px`,
            left: '50%',
            top: '50%',
            transform: `
              translate(
                calc(-50% + ${rotatedX}px),
                calc(-50% + ${rotatedY}px)
              )
              rotate(${angle + chairAngle}deg)
            `,
            background: isUnavailable 
              ? 'var(--mantine-color-dark-5)'
              : 'var(--mantine-color-dark-4)',
            borderRadius: '4px',
            transition: 'all 0.2s ease',
            opacity: isUnavailable ? 0.5 : 0.8,
            boxShadow: isUnavailable 
              ? 'none' 
              : '0 2px 4px rgba(0, 0, 0, 0.2)',
          }}
        />
      );
    });
    
    return chairs;
  };

  return (
    <div
      style={{
        position: 'relative',
        width: '100%',
        height: '600px',
        background: 'var(--mantine-color-dark-8)',
        borderRadius: 'var(--mantine-radius-md)',
        boxShadow: 'inset 0 0 30px rgba(0, 0, 0, 0.3)',
        overflow: 'hidden',
      }}
    >
      {/* Grid lines for visual reference */}
      <div
        style={{
          position: 'absolute',
          width: '100%',
          height: '100%',
          backgroundImage: `
            linear-gradient(rgba(255, 255, 255, 0.03) 1px, transparent 1px),
            linear-gradient(90deg, rgba(255, 255, 255, 0.03) 1px, transparent 1px)
          `,
          backgroundSize: '50px 50px',
          pointerEvents: 'none',
        }}
      />
      
      {tables.map((table) => {
        const isUnavailable = table.occupied || table.seats_number < (guestsCount || 0);
        const isVip = table.type === 'vip';
        
        // Calculate size based on seats
        const baseSize = 45; // Slightly smaller base size
        const sizeMultiplier = Math.sqrt(table.seats_number / 2);
        const width = baseSize * sizeMultiplier;
        const height = baseSize * 0.6; // More rectangular shape
        
        return (
          <Tooltip
            key={table.number}
            label={
              <div style={{ textAlign: 'center' }}>
                <Text size="sm" fw={500}>Стол {table.number}</Text>
                <Text size="xs" c="dimmed">{table.seats_number} мест • {table.type}</Text>
              </div>
            }
            position="top"
            withArrow
            disabled={isUnavailable}
          >
            <div
              style={{
                position: 'absolute',
                left: `${table.x}%`,
                top: `${table.y}%`,
                width: '0',
                height: '0',
                transform: `
                  translate(-50%, -50%)
                `,
              }}
            >
              {/* Chairs */}
              {renderChairs(width, height, table.seats_number, table.angle, isUnavailable)}
              
              {/* Table */}
              <div
                onClick={() => !isUnavailable && onTableClick(table)}
                style={{
                  position: 'absolute',
                  left: '50%',
                  top: '50%',
                  width: `${width}px`,
                  height: `${height}px`,
                  transform: `
                    translate(-50%, -50%)
                    rotate(${table.angle}deg)
                  `,
                  background: isUnavailable 
                    ? 'var(--mantine-color-dark-6)' 
                    : isVip
                      ? 'linear-gradient(45deg, var(--mantine-color-cyan-filled), var(--mantine-color-indigo-filled))'
                      : 'var(--mantine-color-cyan-filled)',
                  cursor: isUnavailable ? 'not-allowed' : 'pointer',
                  borderRadius: '6px',
                  opacity: isUnavailable ? 0.5 : 1,
                  transition: 'all 0.2s ease',
                  boxShadow: isUnavailable 
                    ? 'none'
                    : '0 4px 12px rgba(0, 0, 0, 0.25)',
                  border: `1px solid ${isUnavailable 
                    ? 'var(--mantine-color-dark-5)'
                    : 'rgba(255, 255, 255, 0.1)'}`,
                  zIndex: 1,
                  '&:hover': !isUnavailable && {
                    transform: `
                      translate(-50%, -50%)
                      rotate(${table.angle}deg)
                      scale(1.05)
                    `,
                    boxShadow: '0 6px 16px rgba(0, 0, 0, 0.3)',
                  },
                  '&:after': {
                    content: '""',
                    position: 'absolute',
                    top: '50%',
                    left: '50%',
                    width: '60%',
                    height: '40%',
                    transform: 'translate(-50%, -50%)',
                    border: `1px solid ${isUnavailable 
                      ? 'var(--mantine-color-dark-4)'
                      : 'rgba(255, 255, 255, 0.2)'}`,
                    borderRadius: '3px',
                    pointerEvents: 'none',
                  },
                }}
              />
            </div>
          </Tooltip>
        );
      })}
    </div>
  );
}

export function BookingPage() {
  const navigate = useNavigate();
  const [startTime, setStartTime] = useState<Date | null>(null);
  const [guestsCount, setGuestsCount] = useState<number | ''>(1);
  const [tables, setTables] = useState<Table[]>([]);
  const [selectedTable, setSelectedTable] = useState<Table | null>(null);
  const [showTables, setShowTables] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleSearch = async () => {
    if (!startTime || !guestsCount) {
      notifications.show({
        title: 'Ошибка',
        message: 'Пожалуйста, заполните все поля',
        color: 'red',
      });
      return;
    }

    setLoading(true);
    try {
      const start = dayjs(startTime).format('YYYY-MM-DDTHH:mm:ss');
      const end = dayjs(startTime).add(2, 'hours').format('YYYY-MM-DDTHH:mm:ss');
      const availableTables = await tablesApi.getTables(start, end);
      setTables(availableTables);
      setShowTables(true);
    } catch (error) {
      notifications.show({
        title: 'Ошибка',
        message: 'Не удалось загрузить доступные столы',
        color: 'red',
      });
    } finally {
      setLoading(false);
    }
  };

  const handleTableClick = (table: Table) => {
    if (!table.occupied && table.seats_number >= (guestsCount || 0)) {
      setSelectedTable(table);
    } else if (table.occupied) {
      notifications.show({
        title: 'Ошибка',
        message: 'Этот стол уже занят',
        color: 'red',
      });
    } else {
      notifications.show({
        title: 'Ошибка',
        message: 'За этим столом недостаточно мест',
        color: 'red',
      });
    }
  };

  const handleBooking = async () => {
    if (!startTime || !selectedTable) return;

    setLoading(true);
    try {
      const start = dayjs(startTime).format('YYYY-MM-DDTHH:mm:ss');
      const end = dayjs(startTime).add(2, 'hours').format('YYYY-MM-DDTHH:mm:ss');
      
      const response = await bookingApi.createBooking({
        table_id: selectedTable.number,
        start_time: start,
        end_time: end,
      });

      if (response === 'success') {
        notifications.show({
          title: 'Успех',
          message: 'Стол успешно забронирован',
          color: 'green',
        });
        navigate('/bookings');
      } else {
        notifications.show({
          title: 'Ошибка',
          message: response === 'time_overlap' 
            ? 'У вас уже есть бронирование на это время'
            : response === 'no_tables_available'
            ? 'Нет доступных столов на это время'
            : 'Не удалось забронировать стол',
          color: 'red',
        });
      }
    } catch (error) {
      notifications.show({
        title: 'Ошибка',
        message: 'Не удалось забронировать стол',
        color: 'red',
      });
    } finally {
      setLoading(false);
    }
  };

  return (
    <AppShell
      header={{ height: 80 }}
      footer={{ height: 60 }}
      padding="md"
    >
      <AppShell.Header>
        <Container size="xl" h="100%">
          <Group justify="flex-end" h="100%" pt="md">
            <Tooltip label="Мои бронирования">
              <ActionIcon
                variant="gradient"
                size="lg"
                aria-label="Мои бронирования"
                gradient={{ from: 'cyan', to: 'indigo' }}
                onClick={() => navigate('/bookings')}
              >
                <IconCalendar size={20} />
              </ActionIcon>
            </Tooltip>
          </Group>
        </Container>
      </AppShell.Header>

      <AppShell.Main>
        <Container size="xl" py={50}>
          <Grid justify="center">
            {showTables ? (
              <>
                <Grid.Col span={8}>
                  <Paper shadow="sm" p="md">
                    <TableLayout 
                      tables={tables} 
                      onTableClick={handleTableClick} 
                      guestsCount={guestsCount}
                    />
                  </Paper>
                </Grid.Col>
                <Grid.Col span={4}>
                  <Paper shadow="sm" p="md">
                    <Stack>
                      <DateTimePicker
                        label="Время начала"
                        placeholder="Выберите дату и время"
                        value={startTime}
                        onChange={setStartTime}
                        minDate={new Date()}
                      />
                      <NumberInput
                        label="Количество гостей"
                        placeholder="Введите количество гостей"
                        min={1}
                        max={10}
                        value={guestsCount}
                        onChange={setGuestsCount}
                      />
                      <Button
                        onClick={handleSearch}
                        loading={loading}
                        disabled={!startTime || !guestsCount}
                      >
                        Обновить поиск
                      </Button>
                    </Stack>
                  </Paper>
                </Grid.Col>
              </>
            ) : (
              <Grid.Col span={{ base: 12, sm: 8, md: 6, lg: 4 }}>
                <Paper shadow="md" p="xl" radius="md" withBorder>
                  <Text size="xl" fw={500} ta="center" mb="xl">
                    Забронировать стол
                  </Text>
                  <Stack>
                    <DateTimePicker
                      label="Время начала"
                      placeholder="Выберите дату и время"
                      value={startTime}
                      onChange={setStartTime}
                      minDate={new Date()}
                    />
                    <NumberInput
                      label="Количество гостей"
                      placeholder="Введите количество гостей"
                      min={1}
                      max={10}
                      value={guestsCount}
                      onChange={setGuestsCount}
                    />
                    <Button
                      onClick={handleSearch}
                      loading={loading}
                      disabled={!startTime || !guestsCount}
                      mt="md"
                    >
                      Найти свободные столы
                    </Button>
                  </Stack>
                </Paper>
              </Grid.Col>
            )}
          </Grid>

          <Modal
            opened={!!selectedTable}
            onClose={() => setSelectedTable(null)}
            title="Детали стола"
          >
            {selectedTable && (
              <Stack>
                <Text>
                  <strong>Номер стола:</strong> {selectedTable.number}
                </Text>
                <Text>
                  <strong>Количество мест:</strong> {selectedTable.seats_number}
                </Text>
                <Text>
                  <strong>Тип:</strong> {selectedTable.type}
                </Text>
                <Group mt="md">
                  <Button onClick={handleBooking} loading={loading}>
                    Забронировать
                  </Button>
                  <Button variant="light" onClick={() => setSelectedTable(null)}>
                    Отмена
                  </Button>
                </Group>
              </Stack>
            )}
          </Modal>
        </Container>
      </AppShell.Main>

      <AppShell.Footer>
        <Footer />
      </AppShell.Footer>
    </AppShell>
  );
} 