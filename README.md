# JSON-Time

This simple library implements a Time that supports multiple JSON rendering
forms to preserve time resolution.

Use the standard time.Time if the exact resolution of time is not critical.
If it doesn't matter whether the JSON timestamp is in seconds, milliseconds,
microseconds or nanoseconds resolution, the standard library will suffice.

If you need to maintain a specific resolution, use JSON-Time. This library
provides different types that match the desired resolution:

* `jsontime.SecRes` for second resolution
* `jsontime.MsRes` for millisecond resolution
* `jsontime.UsRes` for microsecond resolution
* `jsontime.NsRes` for nanosecond resolution

## Parsing

When unmarshalling JSON, the input string can have any resolution. The parsed
time will be rounded to the resolution of the `jsontime` type.

For example, consider the timestamp `"2024-07-22T15:05:52.999999999Z"`. This
will result in the following values:

* `jsontime.SecRes` produces `"2024-07-22T15:05:53Z"`
* `jsontime.MsRes` produces `"2024-07-22T15:05:53.000Z"`
* `jsontime.UsRes` produces `"2024-07-22T15:05:53.000000Z"`
* `jsontime.NsRes` produces `"2024-07-22T15:05:52.999999999Z"`

## Rendering

When marshalling JSON, the output string will have the exact resolution of the
`jsontime` type, regardless of the original format.

## Utilities

The following functions return the current time with the respective type:

* `jsontime.SecResNow()`
* `jsontime.MsResNow()`
* `jsontime.UsResNow()`
* `jsontime.NsResNow()`