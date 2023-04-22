use DevDB;

INSERT INTO Label
VALUES (1, 1, "label_1"), -- For choice 1
        (2, 1, "label_2"), -- For choice 1
        (3,  2, "label_3"), -- For choice 2
        (4, 3, "lavel_4"); -- For choice 3

INSERT INTO Choices
VALUES (1, 1), -- For option 1
        (2, 2), -- For option 2
        (3, 3); -- For option 3

INSERT INTO Options
VALUES (1, 1, "Option1_for_menu1"),
        (2, 1, "Option2_for_menu1"),
        (3, 2, "Option3_for_menu2");

INSERT INTO DiscountedTimePeriod
VALUES (1, "01/01/23", "02/01/23");


INSERT INTO Menu
VALUES (1, "small_menu", 0, "thumbnail_image", 100, 10, 30),
        (2, "full_menu", 1, "thumbnail_image", 200, 0, 50, 100, "large_image")