import TypeWriter from "@/components/ui/login/typewriter";
import LoginCard from "./loginCard";

export default function Login() {
  return (
    <div className="min-h-screen bg-hero">
      <div className="grid min-h-screen grid-cols-3 p-4">
        <div className="flex justify-center items-center col-span-2 p-4">
          <TypeWriter />
        </div>
        <div className="flex jsutify-center items-center p-4">
          <LoginCard />
        </div>
      </div>
    </div>
  );
}
