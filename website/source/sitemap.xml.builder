# Simplified from: https://gist.github.com/tommysundstrom/5756032

xml.instruct!

xml.urlset(
  "xmlns" => "http://www.sitemaps.org/schemas/sitemap/0.9",
  "xmlns:xhtml" => "http://www.w3.org/1999/xhtml"
) do
  sitemap.resources.each do |page|
    next unless page.url =~ /\/$/
    xml.url do
      xml.loc("#{base_url}#{page.url}")
      xml.lastmod(Time.now.strftime("%Y-%m-%d"))
    end
  end
end
 