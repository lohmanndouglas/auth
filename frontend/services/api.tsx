import axios, { AxiosResponse, AxiosRequestConfig } from "axios";
import qs from "qs";

const apiConfig: AxiosRequestConfig = {
  timeout: 15000,
  baseURL: `${process.env.REACT_APP_API_URL}`,
  withCredentials: true,
  headers: {
    common: {
      "Content-Type": "application/json",
      "Accept": "application/json",
    },
  },
  paramsSerializer: params => qs.stringify(params, { indices: false }),
};

const axiosInstance = axios.create(apiConfig);

export default function useApi() {
  // const { getAccessTokenSilently } = useAuth0();

  async function getTokenFromYourAPIServer<T, R = AxiosResponse<T>>(form: FormData, token: string): Promise<R> {
    // const token = "asd";
    return axiosInstance.post("/auth/login", form, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  }

  async function updateProfilePicture<T, R = AxiosResponse<T>>(form: FormData): Promise<R> {
    const token = "asd";
    return axiosInstance.put("/api/users/update/profile/picture", form, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  }


  async function createProperty<T, R = AxiosResponse<T>>(form: FormData): Promise<R> {
    const token = "asd";
    return axiosInstance.post("/api/properties", form, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  }

  return {
    getTokenFromYourAPIServer,
    updateProfilePicture,
    createProperty,
  };
}
