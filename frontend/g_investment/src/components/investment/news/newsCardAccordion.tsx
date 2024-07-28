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
import { NewsStock } from "@/types/newsTypes";
import Scores from "./scores";
import { getCompanyNameOfStocks } from "@/actions/newsDataActions";
import { useQuery } from "@tanstack/react-query";

type NewsCardAccordionProps = {
  stocks: NewsStock[];
};

export default function NewsCardAccordion({ stocks }: NewsCardAccordionProps) {
  const { data: companyNames } = useQuery({
    queryKey: ["companyNames", stocks],
    queryFn: () => getCompanyNameOfStocks(stocks),
  });

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
                      {companyNames && companyNames[stock.Stock.symbol]}
                    </CardDescription>
                  </CardHeader>
                  <CardContent>
                    <div>
                      <p className="mb-3">Relevance Score</p>
                      <Scores score={stock.relevance_score} />
                      <br />
                      <p className="mb-3">Analysis Score</p>
                      <Scores score={stock.stock_sentiment_score} />
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
