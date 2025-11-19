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
