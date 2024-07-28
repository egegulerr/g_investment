import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { NewsTableData } from "@/types/newsTypes";
import { Row } from "@tanstack/react-table";
import NewsCardAccordion from "./newsCardAccordion";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import Scores from "./scores";

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
          <Scores score={row.original.sentimentalAnalysisScore} />
        </DialogHeader>
        <NewsCardAccordion stocks={row.original.NewsStocks} />
        <DialogFooter></DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
