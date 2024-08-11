"use client";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { NewsStock } from "@/types/investmentTypes";
import Score from "./score";
import { getCompanyNameOfStocks } from "@/actions/investmentDataActions";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";

type NewsCardAccordionProps = {
  stocks: NewsStock[];
};

export default function NewsCardStocksAccordion({
  stocks,
}: NewsCardAccordionProps) {
  const { data: companyNames, isFetching } = useQuery({
    queryKey: ["companyNames", stocks],
    queryFn: () => getCompanyNameOfStocks(stocks),
  });

  useEffect(() => {
    console.log("News stocks here ", stocks);
  }, [companyNames]);

  return (
    <Accordion type="single" collapsible defaultValue="item-1">
      <AccordionItem value="item-1">
        <AccordionTrigger>Stock Scores</AccordionTrigger>
        <AccordionContent>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 w-full">
            {stocks.map((stock: NewsStock) => {
              return (
                <Card key={stock.ID}>
                  <CardHeader>
                    <CardTitle>{stock.Stock.symbol}</CardTitle>
                    <CardDescription>
                      {isFetching
                        ? "Finding company name..."
                        : companyNames && companyNames[stock.Stock.symbol]}
                    </CardDescription>
                  </CardHeader>
                  <CardContent>
                    <div>
                      <p className="mb-3">Relevance Score</p>
                      <Score score={stock.relevance_score} />
                      <br />
                      <p className="mb-3">Analysis Score</p>
                      <Score score={stock.stock_sentiment_score} />
                      <br />
                    </div>
                  </CardContent>
                  <CardFooter>
                    <span>{stock.stock_sentiment_label}</span>
                  </CardFooter>
                </Card>
              );
            })}
          </div>
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  );
}
