<?php

class HighSchoolSweetheart
{
    public function firstLetter(string $name): string
    {
        $name=trim($name);
        return substr($name,0,1);
    }

    public function initial(string $name): string
    {
        return strtoupper( $this->firstLetter($name) . ".");
  
    }

    public function initials(string $name): string
    { 
        $arr= explode( " ",$name);
        return $this->initial($arr[0] ). " " . $this->initial($arr[1]) ;
    }

    public function pair(string $sweetheart_a, string $sweetheart_b): string
{
 $a = $this->initials($sweetheart_a);
    $b = $this->initials($sweetheart_b);

    $lines = [
        "     ******       ******",
        "   **      **   **      **",
        " **         ** **         **",
        "**            *            **",
        "**                         **",
        "**     " . $a . "  +  " . $b . "     **",
        " **                       **",
        "   **                   **",
        "     **               **",
        "       **           **",
        "         **       **",
        "           **   **",
        "             ***",
        "              *"
    ];

    return implode("\n", $lines);
}
}
 