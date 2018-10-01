# About the NBT Format and Data Structures

The NBT spec (no longer directly available online) describes the format in a way which is curious; if you follow it, you can write code that can read and write the files, but it *makes no sense*.

This is because the specification is fundamentally confused as to what the data structures are. So's Minecraft; this is closely related to why the top-level entry in most NBT files is an empty compound tag with a single named compound tag in it.

## Tag types, Names, and Payloads, Oh My

The NBT spec claims that the basic structure of a tag is type/name/payload. This is just plain wrong. That structure occurs exactly once at the top level of a file, and within compound tags, but it's not the actual representation of things in general. The actual representation of things in general is just their payloads. The tag type is present only if needed; with a List, the tag type is present only once at the head of the list. For ByteArray, IntArray, and LongArray, it's not present at all, but implicit in the container's type. Names are present only for the top-level tag of a file (which usually has an empty name in real data) and in compound tags.

Once you understand this, most of the horrendously ugly code goes away. The correct representation for a Compound tag isn't a map of strings to {type,name,payload} tuples; it's a map of strings to payloads, which innately have types.

This library reads and writes Minecraft data just fine, but the actual data structures used make any kind of sense at all, and thus don't look like the things in the surreally incoherent spec.
