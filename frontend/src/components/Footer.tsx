import { Text, Container, Group } from '@mantine/core';

export function Footer() {
  return (
    <Container size="xl" h="100%">
      <Group justify="center" h="100%" py="md" gap="xs">
        <Text 
          size="sm" 
          c="dimmed"
          style={{
            background: 'linear-gradient(45deg, var(--mantine-color-cyan-filled), var(--mantine-color-indigo-filled))',
            backgroundClip: 'text',
            WebkitBackgroundClip: 'text',
            WebkitTextFillColor: 'transparent',
            fontWeight: 500,
          }}
        >
          POWERED WITH BOOKIT
        </Text>
        <Text size="sm" c="dimmed">•</Text>
        <Text 
          size="sm" 
          c="dimmed"
          style={{ fontWeight: 500 }}
        >
          © Команда ГОСТа/2
        </Text>
      </Group>
    </Container>
  );
} 