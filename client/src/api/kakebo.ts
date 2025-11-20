import { apiClient } from "./index";
export const login = async (email: string, password: string) => {
  const response = await apiClient.post<{ token: string }>("/login", {
    email,
    password,
  });
  return response.data;
};

export const fetchUserFinancialData = async (token: string) => {
  const response = await apiClient.get("/user/top", {
    headers: {
      Authorization: `Bearer ${token} `,
    },
  });
  return response.data;
};
