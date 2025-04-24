import React, { useState, memo } from "react";
import { LinkIcon, ArrowRightIcon } from "lucide-react";
import { shortenURL } from "../lib/fetch";
import { useMutation, useQuery } from "react-query";
import { queryClient, useApp } from "../App";
import { isValidUrl } from "../helper/helper";

export const URLShortenerForm = () => {
  const [url, setUrl] = useState("");
  const [isShortening, setIsShortening] = useState(false);
  const [error, setError] = useState("");
  const { sessionId } = useApp();

  const shortening = useMutation(async (url) => await shortenURL(url), {
    onMutate: () => {
      setIsShortening(true);
    },
    onSuccess: (data) => {
      setIsShortening(false);
      queryClient.invalidateQueries("urls");
    },
    onError: (err) => {
      setIsShortening(false);
      console.log(err);
    },
  });

  const onSubmit = (e) => {
    e.preventDefault();
    if (isValidUrl(url)) {
      const request = {
        origin_url: url,
        session_id: sessionId,
      };
      shortening.mutate(request);
      setError("");
    } else {
      setError("please enter valid URL!");
    }
  };

  return (
    <div className="bg-blue-950 p-4 my-4 rounded-2xl flex flex-col items-center shadow shadow-gray-500">
      <div className="flex flex-col items-center justify-center gap-4">
        <h2 className="text-4xl text-center">Shorten Your Long Long URL</h2>
        <p className="text-xl text-center">
          Create short and memorable link in second
        </p>
      </div>
      <form onSubmit={onSubmit} className="relative w-full lg:w-1/2 pt-6">
        <div className="flex flex-col md:flex-row gap-4">
          <div className="relative flex-grow">
            <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
              <LinkIcon className="h-5 w-5 text-gray-400" />
            </div>
            <input
              type="text"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder="Paste your long URL here"
              className={`w-full pl-10 pr-4 py-3 rounded-md text-gray-300 focus:ring-2 border border-gray-300 focus:ring-blue-400 focus:outline-none ${
                error ? "border-2 border-red-400" : ""
              }`}
            />
          </div>
          <button
            type="submit"
            disabled={isShortening}
            className={`px-6 py-3 bg-blue-800 text-white font-medium rounded-md hover:bg-blue-900 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 flex items-center justify-center transition-colors ${
              isShortening ? "opacity-75 cursor-not-allowed" : ""
            }`}
          >
            {isShortening ? (
              <svg
                className="animate-spin h-5 w-5 mr-2"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  className="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  strokeWidth="4"
                ></circle>
                <path
                  className="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
            ) : (
              <>
                Shorten <ArrowRightIcon className="ml-2 h-5 w-5" />
              </>
            )}
          </button>
        </div>
        {error && (
          <p className="mt-2 text-white bg-red-500 p-2 rounded text-sm">
            {error}
          </p>
        )}
      </form>
    </div>
  );
};
