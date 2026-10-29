*README.txt*                                ID's plain text format for documents
                                                                   Author: CToID

A basic and simple plain text format for writing documents.

Basic structure and format of the document:

                80 chars (or whatever the title line's width is)
    |<------------------------------------------------------------------------>|
    +--------------------------------------------------------------------------+
    |*Filename.txt*                                                  Title here|
    |                                                             Metadata: foo|
    |                                                            Metadata2: bar|
    |                                                                          |
    |Preface or introduction or text without sections. Preface or introduction |
    |or text without sections.                                                 |
    |                                                                          |
    |Paragraph paragraph paragraph paragraph paragraph paragraph paragraph and,|
    |paragraph.                                                                |
    |                                                                          |
    |Paragraph paragraph paragraph paragraph paragraph paragraph paragraph and,|
    |paragraph.                                                                |
    |                                                                          |
    |==========================================================================|
    |Heading level 1                                                           |
    |                                                                          |
    |Paragraph paragraph paragraph paragraph paragraph paragraph paragraph and,|
    |paragraph.                                                                |
    |                                                                          |
    |Paragraph paragraph paragraph paragraph paragraph paragraph paragraph and,|
    |paragraph.                                                                |
    |                                                                          |
    |    codeblock (indent with exactly 4 spaces)                              |
    |    codeblock (indent with exactly 4 spaces)                              |
    |    codeblock (indent with exactly 4 spaces)                              |
    |                                                                          |
    |Paragraph paragraph paragraph paragraph paragraph paragraph paragraph and,|
    |paragraph.                                                                |
    |                                                                          |
    |--------------------------------------------------------------------------|
    |Heading level 2                                                           |
    |                                                                          |
    |Paragraph paragraph paragraph paragraph paragraph paragraph paragraph and,|
    |paragraph. Something something, see the website. https://www.example.com. |
    |                                                                          |
    |Paragraph paragraph paragraph paragraph paragraph paragraph paragraph and,|
    |paragraph. Something something, see the website[1] or another website[2]. |
    |                                                                          |
    |[1]: https://www.example.com                                              |
    |[2]: https://www.example2.com                                             |
    |                                                                          |
    |..........................................................................|
    |Heading level 3                                                           |
    |                                                                          |
    |Paragraph paragraph paragraph paragraph paragraph paragraph paragraph and,|
    |paragraph some cool image[1]. Note that reference numbering restarts from |
    |1 in each section.                                                        |
    |                                                                          |
    |Anything (except newline) behind `:` in the reference is valid. Leading   |
    |spaces are ignored.                                                       |
    |                                                                          |
    |Things between backticks may be displayed specially, like backticks in    |
    |Markdown. To use backticks in backticks, write two consecutive backticks, |
    |just like escaping a single quote in SQL.                                 |
    |                                                                          |
    |[1]: somepath/with/space/image 1.png                                      |
    +--------------------------------------------------------------------------+

It is fine if the content exceeds the line width, but this should be avoided.

Lists or tables? Just put them into codeblocks.

Emphasize text? Use _whatever_ you want, but not backticks! I personally prefer
underscores over asterisks.

Some kind of conversion tool? Maybe in the future when I need to convert it into
other formats.
