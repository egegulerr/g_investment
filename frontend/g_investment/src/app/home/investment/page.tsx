import { getNewsTableData } from "@/actions/newsDataActions";
import News from "@/components/investment/news/news";
import { TypographyH1 } from "@/components/ui/typographyh1";
import { TypographyP } from "@/components/ui/typographyp";
import {
  QueryClient,
  HydrationBoundary,
  dehydrate,
  useQueryClient,
} from "@tanstack/react-query";

export default async function Investment() {
  const queryClient = new QueryClient();
  await queryClient.prefetchQuery({
    queryKey: ["news"],
    queryFn: getNewsTableData,
  });

  return (
    <div>
      <TypographyH1>News</TypographyH1>
      <TypographyP>Here is a list of your news today!</TypographyP>
      <HydrationBoundary state={dehydrate(queryClient)}>
        <News />
      </HydrationBoundary>
    </div>
  );
}
