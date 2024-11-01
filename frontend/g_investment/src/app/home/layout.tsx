import Navbar from "@/components/navbar/navbar";
import React from "react";

export default function HomeLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="min-h-screen">
      <div className="p-8">
        <Navbar />
        <main className="p-16">{children}</main>
      </div>
    </div>
  );
}
