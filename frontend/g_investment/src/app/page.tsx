import { NextUIProvider, Spinner } from "@nextui-org/react";
import * as React from "react";

export default function Home() {
  return (
    <NextUIProvider>
      <div>Hello it is base page</div>
      <div className="flex gap-4">
        <Spinner color="default" />
        <Spinner color="primary" />
        <Spinner color="secondary" />
        <Spinner color="success" />
        <Spinner color="warning" />
        <Spinner color="danger" />
      </div>
    </NextUIProvider>
  );
}
