*README.txt*                                ID's plain text format for documents
                                                                   Author: CToID

A basic and simple plain text format for writing documents.

Basic structure and format of the document:

                                      80 chars
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
    |Any thing (space included) behind `:` in the reference is valid.          |
    |                                                                          |
    |[1]: somepath/with/space/image 1.png                                      |
    +--------------------------------------------------------------------------+

Lists or tables? Just put them into codeblocks.

Emphasize text? Use _whatever_ you want, but not backticks! I personally prefer
underscores over asterisks.

Some kind of conversion tool? Maybe in the future when I need to convert it into
other formats.
