"use client";
import { useQuery } from "@tanstack/react-query";
import NewsTable from "../table/newsTable";
import { getNewsTableData } from "@/actions/investmentDataActions";
import { NewsTableData } from "@/types/investmentTypes";
import { columns } from "../table/columns";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import Stocks from "../stocks/stocks";

export default function News() {
  const [isGrouped, setIsGrouped] = useState(false);
  const { data } = useQuery<NewsTableData[]>({
    queryKey: ["news"],
    queryFn: getNewsTableData,
  });

  return (
    <div className="flex flex-col">
      <Button className="self-end" onClick={() => setIsGrouped(!isGrouped)}>
        {isGrouped ? "Show Table" : "Group by Stocks"}
      </Button>
      {isGrouped ? (
        <Stocks />
      ) : (
        <NewsTable columns={columns} data={data as NewsTableData[]} />
      )}
    </div>
  );
}
