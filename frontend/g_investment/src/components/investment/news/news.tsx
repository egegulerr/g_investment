"use client";
import { useQuery } from "@tanstack/react-query";
import NewsTable from "../table/newsTable";
import { getNewsTableData } from "@/actions/newsDataActions";
import { useEffect } from "react";
import { NewsTableData } from "@/types/newsTypes";
import { columns } from "../table/columns";

export default function News() {
  const { data, error, isFetching } = useQuery<NewsTableData[]>({
    queryKey: ["news"],
    queryFn: getNewsTableData,
  });

  useEffect(() => {
    console.log("data is here", data);
  }, [data]);

  return <NewsTable columns={columns} data={data as NewsTableData[]} />;
}
