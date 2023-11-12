type LoginRequest = { username: string; password: string; otp?: string };
type RegisterRequest = { phone: string; password: string; title: string, confirmPassword: string };

type LoginResponse = {
  jwtToken: string;
  otp?: boolean;
};

type RegisterResponse = {
  token: string;
};