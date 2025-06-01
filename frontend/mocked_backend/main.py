from fastapi import FastAPI, HTTPException, Header, Depends, Request
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from typing import List, Optional
from datetime import datetime, timedelta
import uuid

app = FastAPI()

# Enable CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"], 
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
    expose_headers=["*"],
)

# Mock database
users = {}  # phone_number -> {name, token}
bookings = []  # List of bookings
mock_tables = [
    # Окна (столики на 2-4 персоны)
    {
        "number": 1,
        "x": 15,
        "y": 15,
        "angle": 0,
        "type": "window",
        "seats_number": 2,
        "occupied": False
    },
    {
        "number": 2,
        "x": 15,
        "y": 35,
        "angle": 0,
        "type": "window",
        "seats_number": 2,
        "occupied": False
    },
    {
        "number": 3,
        "x": 15,
        "y": 55,
        "angle": 0,
        "type": "window",
        "seats_number": 4,
        "occupied": False
    },
    {
        "number": 4,
        "x": 15,
        "y": 75,
        "angle": 0,
        "type": "window",
        "seats_number": 4,
        "occupied": False
    },
    # Центр зала (столики на 4-6 персон)
    {
        "number": 5,
        "x": 40,
        "y": 25,
        "angle": 0,
        "type": "standard",
        "seats_number": 4,
        "occupied": False
    },
    {
        "number": 6,
        "x": 40,
        "y": 50,
        "angle": 0,
        "type": "standard",
        "seats_number": 6,
        "occupied": False
    },
    {
        "number": 7,
        "x": 40,
        "y": 75,
        "angle": 0,
        "type": "standard",
        "seats_number": 4,
        "occupied": False
    },
    # VIP-зона (большие столы)
    {
        "number": 8,
        "x": 65,
        "y": 20,
        "angle": 0,
        "type": "vip",
        "seats_number": 8,
        "occupied": False
    },
    {
        "number": 9,
        "x": 65,
        "y": 50,
        "angle": 0,
        "type": "vip",
        "seats_number": 6,
        "occupied": False
    },
    {
        "number": 10,
        "x": 65,
        "y": 80,
        "angle": 0,
        "type": "vip",
        "seats_number": 8,
        "occupied": False
    },
    # Барная зона (высокие столики)
    {
        "number": 11,
        "x": 90,
        "y": 15,
        "angle": 45,
        "type": "bar",
        "seats_number": 2,
        "occupied": False
    },
    {
        "number": 12,
        "x": 90,
        "y": 30,
        "angle": 45,
        "type": "bar",
        "seats_number": 2,
        "occupied": False
    },
    {
        "number": 13,
        "x": 90,
        "y": 45,
        "angle": 45,
        "type": "bar",
        "seats_number": 2,
        "occupied": False
    },
    # Уютные угловые места (диагональное расположение)
    {
        "number": 14,
        "x": 85,
        "y": 70,
        "angle": -45,
        "type": "corner",
        "seats_number": 4,
        "occupied": False
    },
    {
        "number": 15,
        "x": 85,
        "y": 90,
        "angle": -45,
        "type": "corner",
        "seats_number": 4,
        "occupied": False
    },
]

# Models
class AuthRequest(BaseModel):
    name: str
    phone_number: str

class AuthResponse(BaseModel):
    access_token: str

class BookingRequest(BaseModel):
    start_time: str
    end_time: str

# Helper functions
def get_current_user(x_auth_token: str = Header(None)) -> dict:
    if not x_auth_token:
        raise HTTPException(status_code=403, detail="Not authenticated")
    
    for phone_number, user_data in users.items():
        if user_data.get("token") == x_auth_token:
            return {"phone_number": phone_number, "name": user_data["name"]}
    
    raise HTTPException(status_code=403, detail="Invalid token")

@app.middleware("http")
async def log_requests(request: Request, call_next):
    print("=== Request Headers ===")
    for header, value in request.headers.items():
        print(f"{header}: {value}")
    print("=====================")
    response = await call_next(request)
    return response

# Auth endpoint
@app.post("/auth/token", response_model=AuthResponse)
async def login(auth_data: AuthRequest):
    # Generate a simple token
    token = str(uuid.uuid4())
    users[auth_data.phone_number] = {
        "name": auth_data.name,
        "token": token
    }
    return {"access_token": token}

# Booking endpoints
@app.get("/booking/info")
async def get_bookings(current_user: dict = Depends(get_current_user)):
    user_bookings = [
        booking for booking in bookings 
        if booking["user_phone"] == current_user["phone_number"]
    ]
    return {"list": user_bookings}

@app.delete("/booking")
async def delete_booking(booking_id: str, current_user: dict = Depends(get_current_user)):
    booking_idx = None
    for idx, booking in enumerate(bookings):
        if (booking["booking_id"] == booking_id and 
            booking["user_phone"] == current_user["phone_number"]):
            booking_idx = idx
            break
    
    if booking_idx is not None:
        bookings.pop(booking_idx)
        return {"status": "success"}
    
    raise HTTPException(status_code=404, detail="Booking not found")

@app.get("/tables")
async def get_tables(
    start: str,
    end: str,
    current_user: dict = Depends(get_current_user)
):
    # Convert string times to datetime
    start_time = datetime.fromisoformat(start)
    end_time = datetime.fromisoformat(end)
    
    # Check which tables are occupied
    tables = mock_tables.copy()
    for table in tables:
        table["occupied"] = False
        for booking in bookings:
            booking_start = datetime.fromisoformat(booking["start_time"])
            booking_end = datetime.fromisoformat(booking["end_time"])
            if (
                booking["table_id"] == str(table["number"]) and
                (
                    (start_time <= booking_start <= end_time) or
                    (start_time <= booking_end <= end_time) or
                    (booking_start <= start_time and booking_end >= end_time)
                )
            ):
                table["occupied"] = True
                break
    
    return {"list": tables}

@app.post("/booking")
async def create_booking(
    booking_data: BookingRequest,
    current_user: dict = Depends(get_current_user)
):
    # Convert booking times to datetime for comparison
    new_start = datetime.fromisoformat(booking_data.start_time)
    new_end = datetime.fromisoformat(booking_data.end_time)

    # Check if user has any overlapping bookings
    for booking in bookings:
        if booking["user_phone"] == current_user["phone_number"]:
            booking_start = datetime.fromisoformat(booking["start_time"])
            booking_end = datetime.fromisoformat(booking["end_time"])
            
            if (
                (new_start <= booking_start <= new_end) or
                (new_start <= booking_end <= new_end) or
                (booking_start <= new_start and booking_end >= new_end)
            ):
                return {"status": "time_overlap"}
    
    # Find first available table
    available_table = None
    for table in mock_tables:
        table_is_available = True
        # Check if table is booked for the requested time
        for booking in bookings:
            booking_start = datetime.fromisoformat(booking["start_time"])
            booking_end = datetime.fromisoformat(booking["end_time"])
            if (
                booking["table_id"] == str(table["number"]) and
                (
                    (new_start <= booking_start <= new_end) or
                    (new_start <= booking_end <= new_end) or
                    (booking_start <= new_start and booking_end >= new_end)
                )
            ):
                table_is_available = False
                break
        
        if table_is_available:
            available_table = table
            break
    
    if not available_table:
        return {"status": "no_tables_available"}
    
    # Create new booking
    new_booking = {
        "booking_id": str(uuid.uuid4()),
        "table_id": str(available_table["number"]),
        "user_phone": current_user["phone_number"],
        "start_time": booking_data.start_time,
        "end_time": booking_data.end_time
    }
    bookings.append(new_booking)
    
    return {"status": "success"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=3000) 