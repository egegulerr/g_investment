import { getCompanyNameOfStock } from "@/actions/investmentDataActions";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { NewsGroupByStocks } from "@/types/investmentTypes";
import { useQuery } from "@tanstack/react-query";

export default function StocksCard({ symbol, news }: NewsGroupByStocks) {
  const { data: companyName, isFetching } = useQuery({
    queryKey: ["companyName", symbol],
    queryFn: () => getCompanyNameOfStock(symbol),
  });

  return (
    <Card>
      <CardHeader>
        <CardTitle>{symbol}</CardTitle>
        <CardDescription>
          {companyName
            ? companyName
            : isFetching
            ? "Finding company name..."
            : ""}
        </CardDescription>
      </CardHeader>
      <CardContent></CardContent>
      <CardFooter></CardFooter>
    </Card>
  );
}
