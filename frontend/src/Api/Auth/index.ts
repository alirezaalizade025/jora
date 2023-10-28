import API from 'utils/Api';

const login = async (props: LoginRequest) => {
  const response = await API.post<LoginResponse>(`/auth/login`, props);
  return response;
};

const logout = async () => {
  const response = await API.post(`/auth/logout`);
  return response;
};

const AuthApi = {
  login,
  logout,
};

export default AuthApi;
