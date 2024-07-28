export interface NewsTableData {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: null;
  title: string;
  authors: string[];
  sentimentalAnalysisScore: number;
  date: Date;
  summary: string;
  image: string;
  url: string;
  NewsStocks: NewsStock[];
}

export interface NewsStock {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: null;
  NewsID: number;
  StockID: number;
  Stock: Stock;
  relevance_score: number;
  stock_sentiment_score: number;
  stock_sentiment_label: StockSentimentLabel;
}

export interface Stock {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: null;
  symbol: string;
}

export enum StockSentimentLabel {
  Bearish = "Bearish",
  Bullish = "Bullish",
  Neutral = "Neutral",
  SomewhatBearish = "Somewhat-Bearish",
  SomewhatBullish = "Somewhat-Bullish",
}
