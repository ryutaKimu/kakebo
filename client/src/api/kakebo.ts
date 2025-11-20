import { apiClient } from "./index";
export const login = async (email: string, password: string) => {
  const response = await apiClient.post<{ token: string }>("/login", {
    email,
    password,
  });
  return response.data;
};

export const fetchUserFinancialData = async () => {
  const response = await apiClient.get("/user/top", {
    withCredentials: true,
  });
  return response.data;
};
