import { useState, useEffect } from 'react';
import { 
  Card, 
  Text, 
  Button, 
  Group, 
  Stack, 
  Container, 
  Title, 
  Modal,
  AppShell,
  ActionIcon,
  Tooltip,
  Divider,
} from '@mantine/core';
import { notifications } from '@mantine/notifications';
import { IconLogout, IconMapPin, IconMail } from '@tabler/icons-react';
import { bookingApi } from '../services/api';
import type { Booking } from '../types/api';
import dayjs from 'dayjs';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { Footer } from '../components/Footer';
import { formatDateTime } from '../utils/dateTime';
import { cafeConfig } from '../config/cafe';

export function BookingsPage() {
  const navigate = useNavigate();
  const { logout } = useAuth();
  const [bookings, setBookings] = useState<Booking[]>([]);
  const [loading, setLoading] = useState(true);
  const [selectedBooking, setSelectedBooking] = useState<Booking | null>(null);

  const fetchBookings = async () => {
    try {
      const data = await bookingApi.getBookings();
      setBookings(data);
    } catch (error) {
      notifications.show({
        title: 'Ошибка',
        message: 'Не удалось загрузить бронирования',
        color: 'red',
      });
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchBookings();
  }, []);

  const handleDelete = async (bookingId: string) => {
    try {
      await bookingApi.deleteBooking(bookingId);
      notifications.show({
        title: 'Успех',
        message: 'Бронирование успешно удалено',
        color: 'green',
      });
      setBookings(bookings.filter(booking => booking.booking_id !== bookingId));
    } catch (error) {
      notifications.show({
        title: 'Ошибка',
        message: 'Не удалось удалить бронирование',
        color: 'red',
      });
    }
  };

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <AppShell
      header={{ height: 80 }}
      footer={{ height: 60 }}
      padding="md"
    >
      <AppShell.Header>
        <Container size="xl" h="100%">
          <Group justify="space-between" h="100%" pt="md">
            <Group>
              <Title order={3}>{cafeConfig.name}</Title>
              <Group gap="xl" c="dimmed">
                <Group gap={4}>
                  <IconMapPin size={14} />
                  <Text size="sm">{cafeConfig.address}</Text>
                </Group>
                <Group gap={4}>
                  <IconMail size={14} />
                  <Text size="sm">{cafeConfig.email}</Text>
                </Group>
              </Group>
            </Group>
            <Group gap="md">
              <Button 
                variant="light"
                onClick={() => navigate('/book')}
                styles={(theme) => ({
                  root: {
                    transition: 'all 0.2s ease',
                    '&:hover': {
                      background: `linear-gradient(45deg, ${theme.colors.cyan[4]}, ${theme.colors.indigo[4]})`,
                      color: 'white',
                    },
                  },
                })}
              >
                Новое бронирование
              </Button>
              <Tooltip label="Выйти">
                <ActionIcon
                  variant="gradient"
                  size="lg"
                  aria-label="Выйти"
                  gradient={{ from: 'red', to: 'pink' }}
                  onClick={handleLogout}
                >
                  <IconLogout size={20} />
                </ActionIcon>
              </Tooltip>
            </Group>
          </Group>
        </Container>
      </AppShell.Header>

      <AppShell.Main>
        <Container size="xl" py={50}>
          <Title mb="xl">Мои бронирования</Title>

          {loading ? (
            <Text>Загрузка бронирований...</Text>
          ) : bookings.length === 0 ? (
            <Text>Бронирования не найдены.</Text>
          ) : (
            <Stack>
              {bookings.map((booking) => (
                <Card key={booking.booking_id} shadow="sm" padding="lg" radius="md" withBorder>
                  <Group justify="space-between" mb="xs">
                    <Text fw={500}>Стол {booking.table_id}</Text>
                    <Group>
                      <Button variant="light" onClick={() => setSelectedBooking(booking)}>
                        Детали
                      </Button>
                      <Button color="red" onClick={() => handleDelete(booking.booking_id)}>
                        Удалить
                      </Button>
                    </Group>
                  </Group>
                  <Text size="sm" c="dimmed">
                    Начало: {formatDateTime(booking.start_time)}
                  </Text>
                  <Text size="sm" c="dimmed">
                    Конец: {formatDateTime(booking.end_time)}
                  </Text>
                </Card>
              ))}
            </Stack>
          )}

          <Modal
            opened={!!selectedBooking}
            onClose={() => setSelectedBooking(null)}
            title={
              <Title order={3} style={{ 
                background: 'linear-gradient(45deg, var(--mantine-color-cyan-filled), var(--mantine-color-indigo-filled))',
                backgroundClip: 'text',
                WebkitBackgroundClip: 'text',
                WebkitTextFillColor: 'transparent',
              }}>
                {cafeConfig.name}
              </Title>
            }
            size="md"
          >
            {selectedBooking && (
              <Stack>
                <Group gap={4}>
                  <IconMapPin size={16} style={{ color: 'var(--mantine-color-dimmed)' }} />
                  <Text size="sm" c="dimmed">{cafeConfig.address}</Text>
                </Group>
                <Group gap={4}>
                  <IconMail size={16} style={{ color: 'var(--mantine-color-dimmed)' }} />
                  <Text size="sm" c="dimmed">{cafeConfig.email}</Text>
                </Group>
                <Divider my="sm" />
                <Text fw={500} size="lg">Детали бронирования</Text>
                <Text>
                  <Text span fw={500}>Стол:</Text> {selectedBooking.table_id}
                </Text>
                <Text>
                  <Text span fw={500}>Время начала:</Text> {formatDateTime(selectedBooking.start_time)}
                </Text>
                <Text>
                  <Text span fw={500}>Время окончания:</Text> {formatDateTime(selectedBooking.end_time)}
                </Text>
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