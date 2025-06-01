import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { notifications } from '@mantine/notifications';
import { authApi, isAuthenticated, logout as apiLogout } from '../services/api';
import type { AuthRequest } from '../types/api';

interface AuthContextType {
  isLoggedIn: boolean;
  login: (data: AuthRequest) => Promise<void>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(isAuthenticated());

  useEffect(() => {
    // Check authentication status on mount and when token changes
    const checkAuth = () => {
      setIsLoggedIn(isAuthenticated());
    };

    // Check initially
    checkAuth();

    // Set up interval to check periodically
    const interval = setInterval(checkAuth, 1000);

    return () => clearInterval(interval);
  }, []);

  const login = async (data: AuthRequest) => {
    try {
      await authApi.login(data);
      setIsLoggedIn(true);
      notifications.show({
        title: 'Успех',
        message: 'Вход выполнен успешно',
        color: 'green',
      });
    } catch (error) {
      notifications.show({
        title: 'Ошибка',
        message: 'Не удалось войти. Попробуйте еще раз.',
        color: 'red',
      });
      throw error;
    }
  };

  const logout = () => {
    apiLogout();
    setIsLoggedIn(false);
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}; 