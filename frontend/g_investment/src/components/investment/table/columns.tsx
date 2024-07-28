"use client";

import { ColumnDef } from "@tanstack/react-table";
import { DataTableColumnHeader } from "./newsTableColumnHeader";
import { NewsStock, NewsTableData } from "@/types/newsTypes";
import { format, parseISO } from "date-fns";

import NewsCard from "../news/newsCard";
import Scores from "../news/scores";

export const columns: ColumnDef<NewsTableData>[] = [
  {
    accessorKey: "url",
    header: "Source",
    cell: ({ row }) => {
      const url = row.getValue("url") as string;
      const hostname = new URL(url).hostname.replaceAll("www.", "");
      return <div>{hostname}</div>;
    },
  },
  {
    accessorKey: "image",
    enableHiding: true,
    header: ({ column }) => {
      return "";
    },
    cell: ({ row }) => {
      return "";
    },
  },
  {
    accessorKey: "title",
    header: "Title",
    cell: ({ row }) => {
      const title = row.getValue("title") as string;
      return <NewsCard row={row}></NewsCard>;
    },
  },
  {
    accessorKey: "NewsStocks",
    header: "Stocks",
    cell: ({ row }) => {
      const stocks = row.getValue("NewsStocks") as NewsStock[];
      return (
        <div>
          {stocks.map((stock) => (
            <div key={stock.ID}>{stock.Stock.symbol}</div>
          ))}
        </div>
      );
    },
  },
  {
    accessorKey: "date",
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Date" />;
    },
    cell: ({ row }) => {
      const dateString = row.getValue("date") as string;
      const date = parseISO(dateString);
      const formattedDate = format(date, "dd/MMMM/yyyy");
      return <span>{formattedDate}</span>;
    },
  },
  {
    accessorKey: "sentimentalAnalysisScore",
    header: ({ column }) => {
      return <DataTableColumnHeader column={column} title="Score" />;
    },
    cell: ({ row }) => {
      return (
        <div>
          <Scores score={row.original.sentimentalAnalysisScore}></Scores>
        </div>
      );
    },
  },
];
