type LoginRequest = { username: string; password: string; otp?: string };
type AdminLoginRequest = { phone: string; password: string };
type RegisterRequest = { phone: string; password: string; title: string, confirmPassword: string };

type AdminLoginResponse = {
  jwtToken: string;
  otp?: boolean;
};

type LoginResponse = {
  jwtToken: string;
  otp?: boolean;
};

type RegisterResponse = {
  jwtToken: string;
};
