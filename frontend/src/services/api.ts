import axios from 'axios';
import Cookies from 'js-cookie';
import type { AuthRequest, AuthResponse, Booking, Table, BookingRequest, BookingStatus } from '../types/api';

const API_URL = '/api';
const TOKEN_COOKIE_NAME = 'auth_token';

// API client setup
const api = axios.create({
  baseURL: API_URL,
  timeout: 5000, // Add timeout
});

// Request interceptor to add auth token
api.interceptors.request.use((config) => {
  const token = Cookies.get(TOKEN_COOKIE_NAME);
  console.log('Token from cookie:', token);
  if (token) {
    config.headers['X-Auth-Token'] = token;
    console.log('Added token to headers:', config.headers);
  }
  console.log('Request:', { url: config.url, method: config.method, data: config.data });
  return config;
});

// Response interceptor to handle 403 errors
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
      
      // Cookie settings based on environment
      const cookieOptions = {
        expires: 30, // 30 days
        sameSite: 'lax' as const,
        path: '/',
        secure: false // отключаем для HTTP
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
    const response = await api.get<{ list: Booking[] }>('/booking/info');
    return response.data.list;
  },

  deleteBooking: async (bookingId: string): Promise<void> => {
    await api.delete(`/booking?booking_id=${bookingId}`);
  },

  createBooking: async (data: BookingRequest): Promise<BookingStatus> => {
    const response = await api.post<{ status: BookingStatus }>('/booking', data);
    return response.data.status;
  },
};

export const tablesApi = {
  getTables: async (start: string, end: string): Promise<Table[]> => {
    const response = await api.get<{ list: Table[] }>('/tables', {
      params: { start, end },
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