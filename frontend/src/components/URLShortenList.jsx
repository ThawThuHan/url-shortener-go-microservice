import { useQuery } from "react-query";
import { getUrls } from "../lib/fetch";
import { useApp } from "../App";
import { useNavigate } from "react-router";
import { CheckIcon, CopyIcon } from "lucide-react";
import { useState } from "react";
import URLCard from "./URLCard";

export default function URLShortenList() {
  const { sessionId, setUrl } = useApp();
  const { data, isLoading, isError, error } = useQuery("urls", () =>
    getUrls(sessionId)
  );

  const result = data?.urls?.sort(
    (a, b) => new Date(b.created_at) - new Date(a.created_at)
  );

  if (isLoading) {
    return (
      <div className="flex justify-center items-center">
        <h3 className="text-center">Loading....</h3>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="flex justify-center items-center">
        <h3 className="text-center text-red-800">Please reload the page.</h3>
      </div>
    );
  }

  return (
    <div>
      <h3 className="text-center">Your Shortened URLs</h3>
      <div className="Container w-full flex flex-col py-2 gap-2">
        {result == undefined ? (
          <div>
            <h3 className="text-center text-gray-600">
              Your Shortened URLs will be appear here!
            </h3>
          </div>
        ) : (
          result.map((url) => {
            return <URLCard key={url.id} url={url} setUrl={setUrl} />;
          })
        )}
      </div>
    </div>
  );
}
