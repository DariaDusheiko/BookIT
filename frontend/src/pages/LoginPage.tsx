import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { TextInput, Button, Paper, Title, Container, Text, Box, rem, AppShell } from '@mantine/core';
import { useAuth } from '../contexts/AuthContext';
import { IconUser, IconPhone } from '@tabler/icons-react';
import { Footer } from '../components/Footer';
import { cafeConfig } from '../config/cafe';

export function LoginPage() {
  const navigate = useNavigate();
  const { login } = useAuth();
  const [name, setName] = useState('');
  const [phoneNumber, setPhoneNumber] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    
    try {
      await login({
        name,
        phone_number: phoneNumber,
      });
      navigate('/bookings');
    } catch (error) {
      console.error('Login failed:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <AppShell
      footer={{ height: 60 }}
      padding={0}
    >
      <AppShell.Main>
        <Box
          h="100vh"
          display="flex"
          style={{
            alignItems: 'center',
            justifyContent: 'center',
            background: 'var(--mantine-color-dark-8)',
            padding: rem(20),
          }}
        >
          <Container size={420} mt="-10vh">
            <Box ta="center" style={{ transform: 'scale(1.2)' }} mb={50}>
              <Title
                size={rem(56)}
                fw={900}
                style={{
                  background: 'linear-gradient(-45deg, #00ffff, #4dabf7, #7048e8, #00ffff)',
                  backgroundSize: '300% 300%',
                  backgroundClip: 'text',
                  WebkitBackgroundClip: 'text',
                  WebkitTextFillColor: 'transparent',
                  textShadow: '0 0 30px rgba(0, 255, 255, 0.3)',
                  animation: 'gradient 10s ease infinite',
                }}
              >
                {cafeConfig.name}
              </Title>
              <Text 
                c="dimmed" 
                mt="xs" 
                size="lg"
                style={{ letterSpacing: rem(2) }}
              >
                БРОНИРОВАТЬ ПРОСТО
              </Text>
            </Box>

            <Paper
              withBorder
              shadow="xl"
              p={30}
              radius="md"
              bg="dark.6"
              style={{
                backdropFilter: 'blur(16px)',
                border: '1px solid var(--mantine-color-dark-4)',
              }}
            >
              <Text size="lg" fw={500} mb="xl" ta="center" c="dimmed">
                С возвращением
              </Text>
              
              <form onSubmit={handleSubmit}>
                <TextInput
                  label="Имя"
                  placeholder="Ваше имя"
                  required
                  size="md"
                  mb="md"
                  leftSection={<IconUser size={16} />}
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                />
                <TextInput
                  label="Номер телефона"
                  placeholder="+7XXXXXXXXXX"
                  required
                  size="md"
                  mb="xl"
                  leftSection={<IconPhone size={16} />}
                  value={phoneNumber}
                  onChange={(e) => setPhoneNumber(e.target.value)}
                />
                <Button 
                  fullWidth 
                  type="submit" 
                  loading={loading}
                  size="md"
                  variant="gradient"
                  gradient={{ from: 'cyan', to: 'indigo' }}
                  style={{
                    transition: 'transform 0.2s ease',
                  }}
                >
                  Войти
                </Button>
              </form>
            </Paper>
          </Container>

          <style>
            {`
              @keyframes gradient {
                0% { background-position: 0% 50%; }
                50% { background-position: 100% 50%; }
                100% { background-position: 0% 50%; }
              }
            `}
          </style>
        </Box>
      </AppShell.Main>

      <AppShell.Footer>
        <Footer />
      </AppShell.Footer>
    </AppShell>
  );
} 