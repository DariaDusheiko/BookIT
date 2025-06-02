export type AuthRequest = {
  name: string;
  phone_number: string;
}

export type AuthResponse = {
  access_token: string;
}

export type Booking = {
  booking_id: string;
  table_id: string;
  start_time: string;
  end_time: string;
}

export type Table = {
  number: number;
  x: number;
  y: number;
  angle: number;
  type: string;
  seats_number: number;
  occupied: boolean;
}

export type BookingRequest = {
  table_id: number;
  start_time: string;
  end_time: string;
}

export type BookingStatus = 'success' | 'user_has_booking' | 'table_booked'; 