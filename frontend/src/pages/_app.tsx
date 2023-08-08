import { AppProps } from 'next/app';
import Head from 'next/head';
import React from 'react';
import './global.css';

declare global {
  interface Window {
    adsbygoogle: any;
  }
}

function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        {process.browser && (
          // google analytics
          <>
            <script
              async
              src="https://www.googletagmanager.com/gtag/js?id=UA-116967778-8"
            ></script>
            <script
              dangerouslySetInnerHTML={{
                __html: `window.dataLayer = window.dataLayer || [];
              function gtag(){dataLayer.push(arguments);}
              gtag('js', new Date());
              gtag('config', 'UA-116967778-8');`,
              }}
            ></script>
          </>
        )}
        {process.browser && (
          // adsense
          <script
            data-ad-client="ca-pub-7134126650568891"
            async
            src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js"
          ></script>
        )}
        <link rel="icon" href="/favicon.ico" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta name="description" content="Generate pairwise testcases online" />
        <title>Pairwise Pict Online</title>
      </Head>
      <Component {...pageProps} />
    </>
  );
}

export default App;
