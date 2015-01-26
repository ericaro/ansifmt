# Markdown support


AnsiRenderer is just a renderer for blackfriday mardown processing.

Targetting the terminal, it cannot fulfill all the markdown rendering.

From the [markdown reference](http://daringfireball.net/projects/markdown/syntax) 


## Block Elements

### Paragraphs and Linebreak

Paragraphs are indented by 4 spaces, long lines are automatically splitted, using the terminal actual width.

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec eget elit dui. Duis pulvinar, dui at lobortis feugiat, sem metus vestibulum erat, in scelerisque sem urna vel metus. Vivamus augue justo, iaculis vel gravida nec, ultricies quis enim. Nunc id sem ullamcorper, hendrerit urna vel, lacinia turpis. Curabitur vulputate porttitor pulvinar. Integer eget orci scelerisque libero eleifend mattis vel sit amet magna. Praesent pretium augue velit, sit amet sollicitudin nulla ullamcorper et. Mauris ullamcorper elit erat, non commodo turpis dignissim ornare. Sed imperdiet venenatis ante. Nam non ante condimentum, congue mauris consequat, gravida nisl. Nam viverra pharetra vulputate. Mauris varius placerat tortor eget auctor. Cras maximus fermentum diam.

Quisque varius vulputate felis, at lobortis velit dignissim at. Aliquam erat volutpat. Maecenas nec sollicitudin elit. Vestibulum a interdum nunc. Praesent volutpat enim et risus facilisis, a interdum ex blandit. Mauris vestibulum sodales nunc in volutpat. Praesent vitae sodales quam, ac euismod purus. Pellentesque consectetur, tortor vel semper tristique, ipsum nunc pretium nunc, sit amet sollicitudin justo leo eu dolor. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.


### Lists

We fully support lists, ordered, and non ordered. We support nesting list.

Unordered List uses the 'bullet' unicode symbol.

  1. First
  2. second
    1. second.one
    2. second.two
      - nest with *simple* emphasis 
      - nest with **double**  emphasis
      - nest with ***triple***  emphasis
      - nest with ~~strikethrough~~
      - nest with nothing special at all
  3. third


### Headers

Headers are supported. Text is Uppercased, and h1 is not indented, h2 as a one space increment, etc.

All other form of text is incremented by 4 spaces.

The current document uses several headers

### Blockquotes

We don't support block quote, mainly because they introduce recursivity (a blockquote contains paragraph etc.) We do not intend to support recursion, because, using terminal output we don't have this luxury. Nevertheless a limited (one level) support for blockquotes is a good idea.

### Code Blocks

Code block are formatted using a fainted font weight. There is no support for "syntax highlight" so far. Nevertheless this could be a good idea.

    func main(){
        fmt.Printf("Hello World!\n")
    }

## Horizontal Rules

We place an horizontally-centered paragraph-separator: `──*──`.

See 'Paragraph and line break' for an example.

## Span Elements
### Links

We display links "label (url)" all beeing underlined.

You can read this readme [online](http://github.com/ericaro/ansifmt/ansiblackfriday)


### Emphasis and Strikethrough

We fully support *simple highlight* as well as  **double highlight** and even  ***triple highlight***, we also support the optional ~~strikethrough~~ (`~~strikethrough~~`)

### Code

Inline code is display just like code blocks. `func String() string {}`


### Images

Images replaced by the `alt` text.


