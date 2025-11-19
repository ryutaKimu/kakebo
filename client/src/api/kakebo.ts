import { apiClient } from "./index";
export const login = async (email: string, password: string) => {
  const response = await apiClient.post("/login", {
    email,
    password,
  });
  const token = response.data.token;
  localStorage.setItem("access_token", token);

  return token;
};

export const fetchUserFinancialData = async (token: string) => {
  const response = await apiClient.get("/user/top", {
    headers: {
      Authorization: `Bearer ${token} `,
    },
  });
  return response.data;
};
