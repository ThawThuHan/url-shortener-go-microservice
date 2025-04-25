import { createBrowserRouter, RouterProvider } from "react-router";
import Home from "./pages/Home";
import { QueryClient, QueryClientProvider } from "react-query";
import { createContext, useContext, useEffect, useState } from "react";
import { v4 as uuidv4 } from "uuid";
import Details from "./pages/Details";
import Redirect from "./pages/Redirect";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/:short_code/details",
    element: <Details />,
  },
  {
    path: "/:short_code",
    element: <Redirect />,
  },
]);

export const queryClient = new QueryClient();

const AppContent = createContext();

export const useApp = () => useContext(AppContent);

function App() {
  const [url, setUrl] = useState(null);
  const [globalMsg, setGlobalMsg] = useState("");
  const [sessionId, setSessionId] = useState(() => {
    return localStorage.getItem("session_id") || null;
  });
  useEffect(() => {
    if (!sessionId) {
      const newSessionId = uuidv4();
      localStorage.setItem("session_id", newSessionId);
      setSessionId(newSessionId);
    }
  }, []);

  return (
    <AppContent.Provider
      value={{ globalMsg, setGlobalMsg, sessionId, url, setUrl }}
    >
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </AppContent.Provider>
  );
}

export default App;
