import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faMoneyBillTrendUp } from "@fortawesome/free-solid-svg-icons";
import LoginForm from "@/components/ui/login/loginform";

export default function LoginCard() {
  return (
    <Card className="w-5/6">
      <CardHeader className="items-center">
        <CardTitle className="">
          <div className="flex justfiy-center items-center gap-4">
            <FontAwesomeIcon
              className="h-7"
              icon={faMoneyBillTrendUp}
              size="xs"
            />
            <span>G. Investment</span>
          </div>
        </CardTitle>
        <CardDescription>Login to see your dashboard</CardDescription>
      </CardHeader>
      <CardContent>
        <LoginForm />
      </CardContent>
    </Card>
  );
}
