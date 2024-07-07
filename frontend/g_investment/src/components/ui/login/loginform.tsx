"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import {
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormDescription,
  FormMessage,
  Form,
} from "../form";
import { Spinner } from "@nextui-org/spinner";
import { Button } from "../button";
import { Input } from "../input";
import { submitLoginForm } from "@/actions/loginFormActions";
import { useState } from "react";

const formSchema = z.object({
  username: z
    .string()
    .min(2, { message: "Username must be longer then 2 chraccter" })
    .max(50),
  password: z
    .string()
    .min(2, { message: "Password must be longer then 2 chraccter" })
    .max(10),
});

export default function LoginForm() {
  const [isLoading, setIsLoading] = useState(false);
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    console.log("submitted", values);
    try {
      setIsLoading(true);
      await submitLoginForm(values);
    } finally {
      setIsLoading(false);
    }
    submitLoginForm(values);
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
              <FormDescription>Type your username</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input type="password" {...field} />
              </FormControl>
              <FormDescription>Type your password</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button className="w-full" type="submit" disabled={isLoading}>
          {isLoading ? <Spinner color="danger" /> : "Submit"}
        </Button>
      </form>
    </Form>
  );
}
