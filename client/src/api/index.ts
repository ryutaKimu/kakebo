import axios from "axios";

export const apiClient = axios.create({
  baseURL: "http://localhost:9090/",
  headers: {
    "Content-Type": "application/json",
  },
});
