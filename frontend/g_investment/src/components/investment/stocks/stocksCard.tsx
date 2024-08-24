import { getCompanyNameOfStock } from "@/actions/investmentDataActions";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import StocksCardContext from "@/context/useStocksCardContext";
import { useQuery } from "@tanstack/react-query";
import React from "react";
import StocksCardHeader from "./stocksCardHeader";
import StocksCardContent from "./stocksCardContent";
import StocksCardFooter from "./stocksCardFooter";

type StocksCardProps = {
  header: React.ReactNode;
  content: React.ReactNode;
  footer: React.ReactNode;
  symbol: string;
};

export default function StocksCard({
  header,
  content,
  footer,
  symbol,
}: StocksCardProps) {
  const { data: companyName, isFetching } = useQuery({
    queryKey: ["companyName", symbol],
    queryFn: () => getCompanyNameOfStock(symbol),
  });

  return (
    <StocksCardContext.Provider value={symbol}>
      <Card>
        <CardHeader>
          <CardTitle className="flex gap-4 items-center">{header}</CardTitle>
          <CardDescription>
            {companyName
              ? companyName
              : isFetching
              ? "Finding company name..."
              : ""}
          </CardDescription>
        </CardHeader>
        <CardContent>{content}</CardContent>
        <CardFooter>{footer}</CardFooter>
      </Card>
    </StocksCardContext.Provider>
  );
}

StocksCard.Header = StocksCardHeader;
StocksCard.Content = StocksCardContent;
StocksCard.Footer = StocksCardFooter;
