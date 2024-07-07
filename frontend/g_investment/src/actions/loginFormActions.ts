"use server";

export async function submitLoginForm(formData: any) {
  console.log("we got here", formData["username"], formData["password"]);
  await new Promise((resolve) => setTimeout(resolve, 2000));
}
