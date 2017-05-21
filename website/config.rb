compass_config do |config|
  config.output_style = :compact
end

set :base_url, "https://looks.wtf"

set :tags, data.tags
set :looks, data.looks
set :current_tag, nil

tags.each do |tag|
  proxy "/#{tag}", "/index.html", locals: {
    current_tag: tag,
    looks: looks.select { |look| look["tags"].include?(tag) }
  }
end

page "/sitemap.xml", :layout => false

configure :build do
  activate :minify_html
  activate :minify_css
  activate :minify_javascript
  activate :asset_hash
  activate :relative_assets
end
