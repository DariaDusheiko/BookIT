# BookIt - Table Booking Service

A React-based table booking service that allows users to reserve tables in a restaurant. The application features user authentication, booking management, and an interactive table layout.

## Features

- User authentication (login/registration)
- View and manage bookings
- Interactive table layout
- Real-time table availability
- Responsive design

## Prerequisites

- Node.js (v16 or higher)
- npm (v7 or higher)

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd bookit
```

2. Install dependencies:
```bash
npm install --legacy-peer-deps
```

3. Create a `.env` file in the root directory and add your backend API URL:
```
VITE_API_URL=http://your-backend-url
```

4. Start the development server:
```bash
npm run dev
```

The application will be available at `http://localhost:5173`

## Project Structure

- `/src/components` - Reusable React components
- `/src/pages` - Page components
- `/src/contexts` - React contexts (Auth context)
- `/src/services` - API services and utilities

## API Integration

The application integrates with a backend API that provides the following endpoints:

- `POST /auth/token` - User authentication
- `GET /booking/info` - Retrieve user bookings
- `DELETE /booking` - Delete a booking
- `GET /tables` - Get available tables
- `POST /booking` - Create a new booking

## Technologies Used

- React
- TypeScript
- Vite
- React Router
- Mantine UI
- Axios
- Day.js
- JS-Cookie

## Development

To run the development server with hot reload:

```bash
npm run dev
```

To build for production:

```bash
npm run build
```

To preview the production build:

```bash
npm run preview
```

# BookIt Backend

A FastAPI-based mock backend for the BookIt table reservation system.

## Features

- User authentication with token-based access
- Table management with mock data
- Booking creation and management
- CORS support for frontend integration

## Setup

1. Create a virtual environment (recommended):
```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

2. Install dependencies:
```bash
pip install -r requirements.txt
```

3. Run the server:
```bash
python main.py
```

The server will start at `http://localhost:3000`

## API Endpoints

### Authentication
- `POST /auth/token`
  - Request: `{ "name": string, "phone_number": string }`
  - Response: `{ "access_token": string }`

### Bookings
- `GET /booking/info`
  - Headers: `X-Auth-Token`
  - Response: `{ "list": [Booking] }`

- `DELETE /booking`
  - Headers: `X-Auth-Token`
  - Query: `booking_id`

- `POST /booking`
  - Headers: `X-Auth-Token`
  - Request: `{ "start_time": string, "end_time": string }`
  - Response: `{ "status": "success" | "user_has_booking" | "table_booked" }`

### Tables
- `GET /tables`
  - Headers: `X-Auth-Token`
  - Query: `start`, `end`
  - Response: `{ "list": [Table] }`

## Mock Data

The backend includes mock data for tables with different configurations:
- 6 tables with varying capacities (2-8 seats)
- Mix of standard and VIP tables
- Different positions in the layout

## Development

The backend uses in-memory storage for demonstration purposes. All data is reset when the server restarts.
