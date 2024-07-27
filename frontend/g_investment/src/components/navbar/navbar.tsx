import React from "react";
import Sidebar from "./sidebar";
import { Pyramid } from "lucide-react";
import UserDropDown from "./userDropDown";

export default function Navbar() {
  return (
    <div className="flex justify-between">
      <Sidebar />
      <Pyramid />
      <UserDropDown />
    </div>
  );
}
