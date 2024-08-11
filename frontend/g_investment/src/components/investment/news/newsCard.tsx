import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { NewsStock, NewsTableData } from "@/types/investmentTypes";
import { Row } from "@tanstack/react-table";
import NewsCardStocksAccordion from "./newsCardStockAccordion";
import Score from "./score";
import { Label } from "@/components/ui/label";

function calculateAverageScore(newsStocks: NewsStock[]) {
  return (
    newsStocks.reduce((acc, stock) => acc + stock.stock_sentiment_score, 0) /
    newsStocks.length
  );
}

function calculateOverallRelevanceScore(newsStocks: NewsStock[]) {
  return (
    newsStocks.reduce((acc, stock) => acc + stock.relevance_score, 0) /
    newsStocks.length
  );
}

function LabelAndScoreGroup({
  label,
  score,
}: {
  label: string;
  score: number;
}) {
  return (
    <div className="flex flex-col gap-2">
      <Label>{label}</Label>
      <Score score={score} />
    </div>
  );
}

export default function NewsCard({ row }: { row: Row<NewsTableData> }) {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <span className="hover:underline hover:font-bold">
          {row.original.title}
        </span>
      </DialogTrigger>
      <DialogContent className="news-card-dialog news-card-overflow">
        <DialogHeader>
          <DialogTitle>Summary</DialogTitle>
          <DialogDescription className="margin-top-1">
            {row.original.summary}
          </DialogDescription>
          <div className="flex margin-top-1">
            <LabelAndScoreGroup
              label="Sentimental Analysis of News"
              score={row.original.sentimentalAnalysisScore}
            />
            <LabelAndScoreGroup
              label="Overall Analysis Score of Stocks"
              score={calculateAverageScore(row.original.NewsStocks)}
            />
            <LabelAndScoreGroup
              label="Overall Relevance Score of Stocks"
              score={calculateOverallRelevanceScore(row.original.NewsStocks)}
            />
          </div>
        </DialogHeader>
        <NewsCardStocksAccordion stocks={row.original.NewsStocks} />
        <DialogFooter></DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
