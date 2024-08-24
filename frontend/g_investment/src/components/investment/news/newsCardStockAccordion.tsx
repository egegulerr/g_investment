"use client";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { NewsStock } from "@/types/investmentTypes";
import Score from "./score";
import { getCompanyNameOfStocks } from "@/actions/investmentDataActions";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";
import StocksCard from "../stocks/stocksCard";

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
                <StocksCard
                  key={stock.ID}
                  header={<StocksCard.Header />}
                  content={
                    <StocksCard.Content>
                      <div>
                        <p className="mb-3">Relevance Score</p>
                        <Score score={stock.relevance_score} />
                        <br />
                        <p className="mb-3">Analysis Score</p>
                        <Score score={stock.stock_sentiment_score} />
                        <br />
                      </div>
                    </StocksCard.Content>
                  }
                  footer={
                    <StocksCard.Footer>
                      <span>{stock.stock_sentiment_label}</span>
                    </StocksCard.Footer>
                  }
                  symbol={stock.Stock.symbol}
                ></StocksCard>
              );
            })}
          </div>
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  );
}
