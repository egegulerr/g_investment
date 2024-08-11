"use client";
import { getNewsGroupByStocks } from "@/actions/investmentDataActions";
import { NewsGroupByStocks } from "@/types/investmentTypes";
import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";
import StocksCard from "./stocksCard";
import {
  AccordionContent,
  Accordion,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";

export default function Stocks() {
  const { data } = useQuery<NewsGroupByStocks[]>({
    queryKey: ["stockNews"],
    queryFn: getNewsGroupByStocks,
  });

  useEffect(() => {
    console.log("Stocks here", data);
  }, [data]);

  return (
    <Accordion type="multiple" defaultValue={["1"]}>
      <AccordionItem value="item-1">
        <AccordionTrigger>Stocks</AccordionTrigger>
        <AccordionContent>
          <div className="grid gap-4 p-4 sm:grid-cols-2 md:grid-cols-3 ">
            {data?.map((stock) => (
              <StocksCard
                key={stock.symbol}
                symbol={stock.symbol}
                news={stock.news}
              />
            ))}
          </div>
        </AccordionContent>
      </AccordionItem>
      <AccordionItem value="item-2">
        <AccordionTrigger>Is this accessible?</AccordionTrigger>
        <AccordionContent></AccordionContent>
      </AccordionItem>
    </Accordion>
  );
}
