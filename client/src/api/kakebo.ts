import { apiClient } from "./index";
export const login = async (email: string, password: string) => {
  const response = await apiClient.post<{ message: string }>("/login", {
    email,
    password,
  });
  return response.data;
};

export const fetchUserFinancialData = async () => {
  const response = await apiClient.get("/user/top");
  return response.data;
};

export const createAccount = async (
  name: string,
  email: string,
  password: string
): Promise<void> => {
  await apiClient.post("/signup", {
    name,
    email,
    password,
  });
};
