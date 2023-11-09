import API from 'utils/Api';

const login = async (props: LoginRequest) => {
  const response = await API.post<LoginResponse>(`/login`, props);
  return response;
};

const register = async (props: RegisterRequest) => {
  const response = await API.post<RegisterResponse>(`/panel/register`, props);
  return response;
};

const logout = async () => {
  const response = await API.post(`/auth/logout`);
  return response;
};

const AuthApi = {
  login,
  logout,
  register
};

export default AuthApi;
