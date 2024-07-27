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
} from "../ui/form";
import { Spinner } from "@nextui-org/spinner";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { submitLoginForm } from "@/actions/loginActions";
import { useMutation } from "@tanstack/react-query";
import { Alert, AlertTitle, AlertDescription } from "../ui/alert";
import { AlertCircle } from "lucide-react";
import { useRouter } from "next/navigation";

const formSchema = z.object({
  email: z
    .string()
    .min(2, { message: "Email must be longer then 2 chraccter" })
    .max(50),
  password: z
    .string()
    .min(2, { message: "Password must be longer then 2 chraccter" })
    .max(10),
});

export default function LoginForm() {
  const router = useRouter();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const {
    mutate: doLogin,
    isPending,
    isError,
  } = useMutation({
    mutationFn: submitLoginForm,
    onSuccess: () => {
      router.push("/home");
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    doLogin(values);
  }

  return (
    <>
      {isError && (
        <Alert variant="destructive">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>Error</AlertTitle>
          <AlertDescription>
            Login Failed. Please check your username and password.
          </AlertDescription>
        </Alert>
      )}

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="email"
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
          <Button className="w-full" type="submit" disabled={isPending}>
            {isPending ? <Spinner color="danger" /> : "Submit"}
          </Button>
        </form>
      </Form>
    </>
  );
}
