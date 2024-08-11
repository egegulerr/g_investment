"use server";

import { NewsStock } from "@/types/investmentTypes";
import { cookies } from "next/headers";

export async function getNewsTableData() {
  const url = process.env.BACKEND_SERVER_URL + "news";
  try {
    const response = await fetch(url, {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Cookie: cookies().toString(),
      },
    });
    const data = await response.json();
    return data.data;
  } catch (error) {
    console.log("Error occured while fetching news table data", error);
    return [];
  }
}

export async function getNewsGroupByStocks() {
  const url = process.env.BACKEND_SERVER_URL + "stock-news";
  try {
    const response = await fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Cookie: cookies().toString(),
      },
    });
    const data = await response.json();
    return data.data;
  } catch (error) {
    console.log("Error occured while fetching news grouped by stocks", error);
    throw error;
  }
}

export async function getCompanyNameOfStocks(stocks: NewsStock[]) {
  let companyNames: { [key: string]: string } = {};
  for (const stock of stocks) {
    companyNames[stock.Stock.symbol] = await getCompanyNameOfStock(
      stock.Stock.symbol
    );
  }
  return companyNames;
}

export async function getCompanyNameOfStock(ticker: string) {
  const url = `https://query2.finance.yahoo.com/v1/finance/search?q=${ticker}`;
  try {
    const response = await fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Cookie: cookies().toString(),
      },
    });
    const data = await response.json();
    return data.quotes[0]?.longname || "Unknown Company";
  } catch (error) {
    console.log(
      `Error occurred while fetching company name from symbol ${ticker}`
    );
    return "Unknown Company";
  }
}
