$(function() {
  $('.tags .tag').click(function() {
    $('.tags .tag').removeClass('selected');
    $(this).addClass('selected');
    var tag = $(this).attr('href').replace('#', '');
    $('.look-collection .cell.' + tag).show();
    $('.look-collection .cell:not(.' + tag + ')').hide();
  })
});
