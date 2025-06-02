import axios from 'axios';
import Cookies from 'js-cookie';
import type { AuthRequest, AuthResponse, Booking, Table, BookingRequest, BookingStatus } from '../types/api';

const API_URL = import.meta.env.VITE_API_URL;
const TOKEN_COOKIE_NAME = 'auth_token';

// API client setup
const api = axios.create({
  baseURL: API_URL,
  timeout: 5000,
  withCredentials: true, // <- Добавьте эту строку для включения отправки учетных данных
});

// Request interceptor to add auth token
api.interceptors.request.use((config) => {
  // Skip adding token for auth endpoint
  if (config.url === '/auth/token') {
    return config;
  }

  const token = Cookies.get(TOKEN_COOKIE_NAME);
  console.log('Token from cookie:', token);
  
  if (token) {
    config.headers['X-Auth-Token'] = token;
    console.log('Added token to headers:', config.headers);
  } else {
    console.warn('No auth token found for protected route:', config.url);
  }
  
  console.log('Request:', { url: config.url, method: config.method, data: config.data });
  return config;
});

// Response interceptor to handle errors
api.interceptors.response.use(
  (response) => {
    console.log('Response:', { url: response.config.url, status: response.status, data: response.data });
    return response;
  },
  (error) => {
    console.error('API Error:', {
      url: error.config?.url,
      method: error.config?.method,
      status: error.response?.status,
      data: error.response?.data,
      message: error.message
    });
    
    if (error.response?.status === 403) {
      Cookies.remove(TOKEN_COOKIE_NAME);
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// API methods
export const authApi = {
  login: async (data: AuthRequest): Promise<string> => {
    try {
      console.log('Attempting login with:', data);
      const response = await api.post<AuthResponse>('/auth/token', data);
      const token = response.data.access_token;
      console.log('Login successful, token received');
      
      const cookieOptions = {
        expires: 30,
        sameSite: 'lax' as const,
        path: '/',
        secure: false
      };
      
      Cookies.set(TOKEN_COOKIE_NAME, token, cookieOptions);
      return token;
    } catch (error) {
      console.error('Login failed:', error);
      throw error;
    }
  },
};

export const bookingApi = {
  getBookings: async (): Promise<Booking[]> => {
    try {
      const response = await api.post<{ bookings: Booking[] }>('/booking/info', {
        body: {},
        headers: {
          'X-Auth-Token': Cookies.get(TOKEN_COOKIE_NAME) || '', // Отправляем токен или пустую строку
          'Content-Type': 'application/json'
        }
      });
      
      // Обеспечиваем, что каждый booking имеет все обязательные поля
      return response.data.bookings.map(booking => ({
        booking_id: booking.booking_id || '',
        table_id: booking.table_id || '',
        start_time: booking.start_time || '',
        end_time: booking.end_time || '',
        // Добавьте другие поля по необходимости
      }));
    } catch (error) {
      console.error('Failed to fetch bookings:', error);
      return []; // Возвращаем пустой массив в случае ошибки
    }
  },

  deleteBooking: async (bookingId: string): Promise<void> => {
    await api.delete('/booking/', {
      data: { booking_id: bookingId }, // Тело запроса
      headers: {
        'Content-Type': 'application/json',
        'X-Auth-Token': Cookies.get(TOKEN_COOKIE_NAME) || ''
      }
    });
  },

  createBooking: async (data: BookingRequest): Promise<BookingStatus> => {

    const formatDate = (dateString: string) => {
      if (!dateString) return dateString;
      if (dateString.includes('Z') || dateString.match(/[+-]\d{2}:\d{2}$/)) {
        return dateString;
      }
      return `${dateString}Z`;
    };

    const formattedData = {
      table_id: data.table_id, // Явно включаем table_id
      start_time: formatDate(data.start_time),
      end_time: formatDate(data.end_time),
      // Другие обязательные поля
    };

    console.log('Sending booking data:', formattedData); // Для отладки

    const response = await api.post<{ status: BookingStatus }>('/booking/', formattedData, {
      headers: {
        'Content-Type': 'application/json',
        'X-Auth-Token': Cookies.get(TOKEN_COOKIE_NAME), // Добавляем токен
      }
    });
    
    return response.data.status;
  },
};

export const tablesApi = {
  getTables: async (start: string, end: string): Promise<Table[]> => {
    const response = await api.post<{ list: Table[] }>('/tables/', {
      start: `${start}Z`, // Просто добавляем Z в конец
      end: `${end}Z`      // Просто добавляем Z в конец
    }, {
      headers: {
        'Content-Type': 'application/json',
        'X-Auth-Token': Cookies.get(TOKEN_COOKIE_NAME),
      }
    });
    return response.data.list;
  },
};

export const isAuthenticated = (): boolean => {
  return !!Cookies.get(TOKEN_COOKIE_NAME);
};

export const logout = (): void => {
  Cookies.remove(TOKEN_COOKIE_NAME);
};