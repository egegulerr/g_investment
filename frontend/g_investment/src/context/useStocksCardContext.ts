import { createContext, useContext } from "react";

const StocksCardContext = createContext<string | null>(null);

export function useStocksCardContext() {
  const context = useContext(StocksCardContext);
  if (context === null) {
    throw new Error(
      "useStocksCardContext must be used within a StocksCardProvider"
    );
  }
  return context;
}

export default StocksCardContext;
