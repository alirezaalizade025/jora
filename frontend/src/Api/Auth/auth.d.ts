type LoginRequest = { username: string; password: string; otp?: string };

type LoginResponse = {
  jwtToken: string;
  otp?: boolean;
};
