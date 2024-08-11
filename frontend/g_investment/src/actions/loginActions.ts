"use server";

import { cookies } from "next/headers";

type loginResponse = {
  data: string;
};

export async function submitLoginForm(formData: any) {
  const url = process.env.BACKEND_SERVER_URL + "login";
  const response = await fetch(url, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(formData),
  });
  const data: loginResponse = await response.json();
  cookies().set({
    name: "jwt",
    value: data.data,
    httpOnly: true,
    path: "/",
    maxAge: 60 * 60 * 24 * 365 * 1000,
    expires: new Date(Date.now() + 60 * 60 * 24),
  });
  return data;
}

export async function checkJwtValidity() {
  try {
    const response = await fetch(
      process.env.BACKEND_SERVER_URL + "checkToken",
      {
        method: "GET",
        credentials: "include",
        headers: {
          Cookie: cookies().toString(),
        },
      }
    );
    return true;
  } catch (error) {
    console.log("Jwt is not valid");
    return false;
  }
}
