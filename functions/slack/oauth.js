export async function onRequest(context) {
  const code = new URL(context.request.url).searchParams.get('code');

  const response = await fetch('https://slack.com/api/oauth.access', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: new URLSearchParams({
      client_id: context.env.SLACK_CLIENT_ID,
      client_secret: context.env.SLACK_CLIENT_SECRET,
      code: code,
    }).toString(),
  });

  const body = await response.json();
  if (body.ok) {
    console.log(`Ok, response: ${JSON.stringify(body)}`);
    return new Response('The looks.wtf Slack App has been successfully added!\nGive it a go with the `/look awe` command in Slack.');
  } else {
    console.log(`Not ok, response: ${JSON.stringify(body)}`);
    throw new Error('Not ok response from Slack');
  }
}
