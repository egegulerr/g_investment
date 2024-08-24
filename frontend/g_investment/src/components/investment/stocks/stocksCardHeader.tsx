import StocksCardContext from "@/context/useStocksCardContext";
import { useContext } from "react";
import CompanyLogo from "../news/companyLogo";

export default function StocksCardHeader() {
  const symbol = useContext(StocksCardContext);

  return (
    <>
      {symbol}
      <CompanyLogo symbol={symbol || ""} />
    </>
  );
}
